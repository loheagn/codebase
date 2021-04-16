package main

import (
	"fmt"
	"strings"
)

var allSeq = `access_id_seq
access_token_id_seq
login_source_id_seq
issue_id_seq
pull_request_id_seq
comment_id_seq
version_id_seq
user_id_seq
public_key_id_seq
two_factor_id_seq
two_factor_recovery_code_id_seq
project_id_seq
pipeline_id_seq
pipe_task_id_seq
build_task_id_seq
push_task_id_seq
repository_id_seq
validation_task_id_seq
validation_result_id_seq
deploy_key_id_seq
collaboration_id_seq
upload_id_seq
watch_id_seq
star_id_seq
follow_id_seq
action_id_seq
attachment_id_seq
issue_user_id_seq
label_id_seq
issue_label_id_seq
milestone_id_seq
mirror_id_seq
webhook_id_seq
hook_task_id_seq
release_id_seq
protect_branch_id_seq
protect_branch_whitelist_id_seq
team_id_seq
org_user_id_seq
team_user_id_seq
team_repo_id_seq
notice_id_seq
email_address_id_seq`

func main() {
	seqs := strings.Split(allSeq, "\n")
	for _, s := range seqs {
		fmt.Printf("alter sequence %s restart 100;\n", s)
	}
}
