
const twitchAuthLink = `https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=${process.env.NEXT_PUBLIC_TWITCH_CLIENT_ID}&redirect_uri=${process.env.NEXT_PUBLIC_TWITCH_REDIRECT_URI}&scope=channel%3Amanage%3Apolls+channel%3Aread%3Apolls&state`

export const appLinks = {
  twithAuth: twitchAuthLink,
};
