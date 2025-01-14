import { MICROSOFT_CLIENT_ID } from '@env';
import { OauthLogin } from './oauthCall';
import { sendServiceToken } from './sendServiceToken';

export async function microsoftLogin(
  apiEndpoint: string,
  sessionToken?: string,
) {
  const setToken = (accessToken: string) => {
    if (sessionToken)
      sendServiceToken(apiEndpoint, accessToken, 'microsoft', sessionToken);
    else sendServiceToken(apiEndpoint, accessToken, 'microsoft');
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
