
import { appLinks } from "@utils/links";
import Link from "next/link";
import { Button } from "ui";


export default function Web() {
  return (
    <div>
      <h1>Web</h1>
      <Link href={appLinks}>Login with twitch</Link>
      <Button> <span>Login with twith</span></Button>
    </div>
  );
}
