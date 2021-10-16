const github = require("@actions/github");
const fs = require("fs");
const path = require("path");

async function run() {
  const token = process.env["GHTOKEN"];
  const root = process.env["GITHUB_WORKSPACE"];
  const octokit = github.getOctokit(token);

  octokit.rest.repos.uploadReleaseAsset({
    owner: github.context.repo.owner,
    repo: github.context.repo.repo,
    release_id: github.context.payload.release.id,
    name: "dsgore",
    data: fs.readFileSync(path.join(root, "build/dsgore")),
  });
}

run();
