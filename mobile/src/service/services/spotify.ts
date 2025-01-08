import { deleteToken, saveToken } from "../token";
import { SPOTIFY_CLIENT_ID, SPOTIFY_SECRET } from '@env';
import { OauthLogin } from '../oauth/oauthCall';

export async function spotifyLogin(apiEndpoint: string, email?: string) {
  console.log('Spotify login');
  const setToken = (accessToken: string) => {
    console.log('Spotify token:', accessToken);
    // sendSpotifyToken(apiEndpoint, accessToken);
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

export async function sendSpotifyToken(
  apiEndpoint: string,
  token: string,
  email?: string,
) {
  try {
    let response;
    if (email) {
      response = await fetch(
        `http://${apiEndpoint}:8080/api/github/mobile/token`,
        {
          method: 'POST',
          body: JSON.stringify({ token: token, email: email }),
        },
      );
    } else {
      response = await fetch(
        `http://${apiEndpoint}:8080/api/github/mobile/token`,
        {
          method: 'POST',
          body: JSON.stringify({ token: token }),
        },
      );
    }
    const data = await response.json();
    if (response.status === 200) {
      console.log('API send Spotify Token success');
    }
    deleteToken('token');
    saveToken('token', data.token);
    return true;
  } catch (error) {
    console.error('Error:', error);
    return false;
  }
}
