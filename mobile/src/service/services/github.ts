import { GITHUB_CLIENT_ID, GITHUB_SECRET } from '@env';
import { OauthLogin } from './oauthCall';
import { sendServiceToken } from './sendServiceToken';

export async function githubLogin(apiEndpoint: string, sessionToken?: string) {
  const setToken = (accessToken: string) => {
    if (sessionToken)
      sendServiceToken(apiEndpoint, accessToken, 'github', sessionToken);
    else sendServiceToken(apiEndpoint, accessToken, 'github');
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
