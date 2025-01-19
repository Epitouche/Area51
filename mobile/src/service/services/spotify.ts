import { SPOTIFY_CLIENT_ID, SPOTIFY_SECRET } from '@env';
import { sendServiceToken, OauthLogin } from './services';

export async function spotifyLogin(apiEndpoint: string, sessionToken?: string) {
  const setToken = async (accessToken: string) => {
    if (sessionToken)
      await sendServiceToken(apiEndpoint, accessToken, 'spotify', sessionToken);
    else await sendServiceToken(apiEndpoint, accessToken, 'spotify');
  };
  const config = {
    clientId: SPOTIFY_CLIENT_ID,
    clientSecret: SPOTIFY_SECRET,
    redirectUrl: 'com.area51-epitech://oauthredirect',
    scopes: ['user-read-email', 'playlist-modify-public'],
    serviceConfiguration: {
      authorizationEndpoint: 'https://accounts.spotify.com/authorize',
      tokenEndpoint: 'https://accounts.spotify.com/api/token',
    },
  };

  if (await OauthLogin({ config, setToken })) return true;
  return false;
}
