data "git_repository" "this" {
  count = var.spacelift_commit_sha != null ? 0 : 1
  path  = "${path.module}/.."
}

locals {
  # Commit sha can be retreived from git when called as a remote git module
  commit_sha       = var.spacelift_commit_sha != null ? var.spacelift_commit_sha : data.git_repository.this[0].commit_sha
  short_commit_sha = substr(local.commit_sha, 0, 7)
}
