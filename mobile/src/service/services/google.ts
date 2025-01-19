import { GOOGLE_CLIENT_ID, GOOGLE_MOBILE_ID, GOOGLE_SECRET } from '@env';
import { sendServiceToken, OauthLogin } from './services';

export async function googleLogin(apiEndpoint: string, sessionToken?: string) {
  const setToken = async (accessToken: string) => {
    if (sessionToken)
      await sendServiceToken(apiEndpoint, accessToken, 'google', sessionToken);
    else await sendServiceToken(apiEndpoint, accessToken, 'google');
  };
  const config = {
    issuer: 'https://accounts.google.com',
    clientId: `${GOOGLE_MOBILE_ID}.apps.googleusercontent.com`,
    redirectUrl: `com.googleusercontent.apps.${GOOGLE_MOBILE_ID}:/oauth2redirect/google`,
    scopes: [
      'openid',
      'https://www.googleapis.com/auth/userinfo.email',
      'https://www.googleapis.com/auth/userinfo.profile',
      'https://www.googleapis.com/auth/gmail.readonly',
      'https://www.googleapis.com/auth/gmail.labels',
      'https://www.googleapis.com/auth/gmail.modify',
      'https://www.googleapis.com/auth/gmail.metadata',
      'https://www.googleapis.com/auth/calendar',
      'https://www.googleapis.com/auth/calendar.events',
    ],
  };

  if (await OauthLogin({ config, setToken })) return true;
  return false;
}

('https://www.googleapis.com/auth/calendar.events');
