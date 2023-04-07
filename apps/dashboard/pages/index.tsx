
import { appLinks } from "@utils/links";
import Link from "next/link";


export const getServerSideProps = async ()=> {
  console.log(process.env.NEXT_PUBLIC_AUTH_SERVER_API_URL);

  return {
    props: {}
  }
  
}

export default function Web() {
  
  return (
    <div>
      <h1>Web</h1>
      <Link href={appLinks.twithAuth}>Login with twitch</Link>      
    </div>
  );
}
