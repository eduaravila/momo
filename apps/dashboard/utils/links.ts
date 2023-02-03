
const twitchAuthLink = `https://id.twitch.tv/oauth2/authorize?response_type=code&client_id=${process.env.NEXT_PUBLIC_TWITCH_APPLICATION_CLIEND_ID}&redirect_uri=${process.env.NEXT_PUBLIC_AUTH_SERVER_API_URL}/${process.env.NEXT_PUBLIC_TWITCH_APPLICATION_REDIRECT_PATH}&scope=${process.env.NEXT_PUBLIC_TWITCH_APPLICATION_SCOPES}&claims=${process.env.NEXT_PUBLIC_TWITCH_APPLICATION_CLAIMS}`


export const appLinks = {
  twithAuth: twitchAuthLink,
};
