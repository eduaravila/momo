
import { appLinks } from "@utils/links";
import Link from "next/link";


export default function Web() {
  
  return (
    <div>
      <h1>Web</h1>
      <Link href={appLinks.twithAuth}>Login with twitch</Link>      
    </div>
  );
}
