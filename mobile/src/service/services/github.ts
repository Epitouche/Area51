import { deleteToken, saveToken } from "../token";
import { GITHUB_CLIENT_ID, GITHUB_SECRET } from '@env';
import { OauthLogin } from '../oauth/oauthCall';

export async function githubLogin(apiEndpoint: string, email?: string) {
  const setToken = (accessToken: string) => {
    if (email) sendGithubToken(apiEndpoint, accessToken, email);
    else sendGithubToken(apiEndpoint, accessToken);
  };
  const config = {
    clientId: GITHUB_CLIENT_ID,
    clientSecret: GITHUB_SECRET,
    scopes: ['repo'],
    redirectUrl: 'com.area51-epitech://oauthredirect',
    serviceConfiguration: {
      authorizationEndpoint: 'https://github.com/login/oauth/authorize',
      tokenEndpoint: 'https://github.com/login/oauth/access_token',
    },
  };

  if (await OauthLogin({ config, setToken })) return true;
  return false;
}

export async function sendGithubToken(
  apiEndpoint: string,
  token: string,
  email?: string,
) {
  try {
    let response;
    if (email) {
      response = await fetch(`http://${apiEndpoint}:8080/api/github/mobile/token`, {
        method: 'POST',
        body: JSON.stringify({ 'token': token, 'email': email }),
      });
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
      console.log('API send Github Token success');
    }
    deleteToken('token');
    saveToken('token', data.token);
    return true;
  } catch (error) {
    console.error('Error:', error);
    return false;
  }
}
