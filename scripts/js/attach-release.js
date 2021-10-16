const github = require("@actions/github");
const fs = require("fs");

async function run() {
  const token = process.env["GHTOKEN"];
  const octokit = github.getOctokit(token);

  octokit.rest.repos.uploadReleaseAsset({
    owner: github.context.repo.owner,
    repo: github.context.repo.repo,
    release_id: github.context.payload.release.id,
    name: "dsgore",
    data: fs.readFileSync("build/dsgore"),
  });
}

run();
