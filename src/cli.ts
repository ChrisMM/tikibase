import { mentions } from "./mentions"

async function main() {
  try {
    await mentions()
  } catch (e) {
    console.log(e)
    process.exit(1)
  }
}

main()
