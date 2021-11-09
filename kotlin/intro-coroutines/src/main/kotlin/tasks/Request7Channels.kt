package tasks

import contributors.*
import kotlinx.coroutines.*
import kotlinx.coroutines.channels.Channel

suspend fun loadContributorsChannels(
    service: GitHubService,
    req: RequestData,
    updateResults: suspend (List<User>, completed: Boolean) -> Unit
) {
    coroutineScope {
        val repos = service
            .getOrgRepos(req.org)
            .also { logRepos(req, it) }
            .body() ?: listOf()
        val channel = Channel<List<User>>(repos.size)
        repos.forEach { repo ->
            launch {
                service
                    .getRepoContributors(req.org, repo.name)
                    .also { logUsers(repo, it) }
                    .bodyList()
                    .also { channel.send(it) }
            }
        }
        launch {
            var users = emptyList<User>()
            repeat(repos.size) { index ->
                channel.receive().also {
                    users = (users + it).aggregate()
                    updateResults(users, index == repos.lastIndex)
                }
            }
        }
    }
}
