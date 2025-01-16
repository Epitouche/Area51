import { GOOGLE_CLIENT_ID, GOOGLE_MOBILE_ID, GOOGLE_SECRET } from '@env';
import { sendServiceToken, OauthLogin } from './services';

export async function googleLogin(apiEndpoint: string, sessionToken?: string) {
  const setToken = (accessToken: string) => {
    if (sessionToken)
      sendServiceToken(apiEndpoint, accessToken, 'google', sessionToken);
    else sendServiceToken(apiEndpoint, accessToken, 'google');
  };
  console.log('googleLogin', GOOGLE_MOBILE_ID);
  const config = {
    issuer: 'https://accounts.google.com',
    clientId: `${GOOGLE_MOBILE_ID}.apps.googleusercontent.com`,
    redirectUrl: `com.googleusercontent.apps.${GOOGLE_MOBILE_ID}:/oauth2redirect/google`,
    // redirectUrl: 'com.area51-epitech://oauthredirect',
    scopes: ['https://mail.google.com/', 'profile', 'email'],
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
