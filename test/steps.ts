import { Before, setWorldConstructor, Given, Then, When } from "cucumber"
import { promises as fsp } from "fs"
import fse from "fs-extra"
import path from "path"
import { mentions } from "../src/mentions"
import assert from "assert"

setWorldConstructor(function CustomWorld() {
  this.workspacePath = function(filePath: string) {
    return path.join(this.workspace, filePath)
  }
})

Before(async function() {
  this.workspace = "./tmp"
  await fse.emptyDir(this.workspace)
})

Given("the workspace contains file {string} with content:", async function(
  filename,
  content
) {
  return fsp.writeFile(this.workspacePath(filename), content)
})

When("running Mentions", async function() {
  await mentions()
})

Then(
  "the workspace should contain the file {string} with content:",
  async function(filename, expectedContent) {
    const actualContent = await fsp.readFile(
      this.workspacePath(filename),
      "utf8"
    )
    assert.equal(actualContent, expectedContent)
  }
)
