import { deleteToken, saveToken } from "../token";
import { MICROSOFT_CLIENT_ID } from '@env';
import { OauthLogin } from '../oauth/oauthCall';

export async function microsoftLogin(apiEndpoint: string, email?: string) {
  const setToken = (accessToken: string) => {
    // if (email) sendGithubToken(apiEndpoint, accessToken, email);
    // else sendGithubToken(apiEndpoint, accessToken);
    console.log('Microsoft token:', accessToken);
  };
  const config = {
    clientId: MICROSOFT_CLIENT_ID,
    scopes: ['Mail.ReadWrite', 'User.Read', 'Mail.Send', 'offline_access'],
    redirectUrl: 'com.area51-epitech://oauthredirect',
    serviceConfiguration: {
      authorizationEndpoint:
        'https://login.microsoftonline.com/common/oauth2/v2.0/authorize',
      tokenEndpoint:
        'https://login.microsoftonline.com/common/oauth2/v2.0/token',
    },
  };

  if (await OauthLogin({ config, setToken })) return true;
  return false;
}

export async function sendMicrosoftToken(
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
