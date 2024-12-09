import { authorize } from 'react-native-app-auth';
import { GITHUB_CLIENT_ID, GITHUB_SECRET } from '@env';
import { saveToken } from '../token';

export const githubLogin = async (setToken: (token: string) => void) => {
  try {
    const authState = await authorize({
      clientId: GITHUB_CLIENT_ID,
      clientSecret: GITHUB_SECRET,
      scopes: ['repo'],
      redirectUrl: 'com.area51-epitech://oauthredirect',
      serviceConfiguration: {
        authorizationEndpoint: 'https://github.com/login/oauth/authorize',
        tokenEndpoint: 'https://github.com/login/oauth/access_token',
      },
    });
    setToken(authState.accessToken);
    saveToken('github', authState.accessToken);
    return true;
  } catch (error) {
    console.error(error);
    setToken('');
    return false;
  }
};
