import { GOOGLE_CLIENT_ID, GOOGLE_MOBILE_ID, GOOGLE_SECRET } from '@env';
import { sendServiceToken, OauthLogin } from './services';

export async function googleLogin(apiEndpoint: string, sessionToken?: string) {
  const setToken = (accessToken: string) => {
    if (sessionToken)
      sendServiceToken(apiEndpoint, accessToken, 'google', sessionToken);
    else sendServiceToken(apiEndpoint, accessToken, 'google');
  };
  const config = {
    issuer: 'https://accounts.google.com',
    clientId: `${GOOGLE_MOBILE_ID}.apps.googleusercontent.com`,
    redirectUrl: `com.googleusercontent.apps.${GOOGLE_MOBILE_ID}:/oauth2redirect/google`,
    scopes: [
      'openid https://www.googleapis.com/auth/userinfo.email',
      'https://www.googleapis.com/auth/userinfo.profile',
      'https://www.googleapis.com/auth/gmail.readonly',
      'https://www.googleapis.com/auth/gmail.labels',
      'https://www.googleapis.com/auth/gmail.modify',
      'https://www.googleapis.com/auth/gmail.metadata',
    ],
  };

  if (await OauthLogin({ config, setToken })) return true;
  return false;
}

// clientId: GITHUB_CLIENT_ID,
// clientSecret: GITHUB_SECRET,
// scopes: ['repo'],
// serviceConfiguration: {
//   authorizationEndpoint: 'https://github.com/login/oauth/authorize',
//   tokenEndpoint: 'https://github.com/login/oauth/access_token',
// },
