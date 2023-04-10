
import { appLinks } from "@utils/links";
import Link from "next/link";


export const getServerSideProps = async ()=> {
  console.log(process.env);

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
