import { GITHUB_CLIENT_ID, GITHUB_SECRET } from '@env';
import { saveToken } from '../token';
import { sendGithubToken } from '../github';
import { OauthLogin } from './oauthCall';

export async function githubLogin(apiEndpoint: string) {
  const setToken = (accessToken: string) => {
    saveToken('github', accessToken);
    sendGithubToken(apiEndpoint, accessToken);
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
