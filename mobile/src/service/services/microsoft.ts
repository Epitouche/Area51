import { MICROSOFT_CLIENT_ID } from '@env';
import { sendServiceToken, OauthLogin } from './services';

export async function microsoftLogin(
  apiEndpoint: string,
  sessionToken?: string,
) {
  const setToken = async (accessToken: string) => {
    if (sessionToken)
      await sendServiceToken(
        apiEndpoint,
        accessToken,
        'microsoft',
        sessionToken,
      );
    else await sendServiceToken(apiEndpoint, accessToken, 'microsoft');
  };
  const config = {
    clientId: MICROSOFT_CLIENT_ID,
    scopes: [
      'openid',
      'profile',
      'Calendars.Read',
      'Calendars.ReadWrite',
      'Calendars.ReadWrite.Shared',
      'Calendars.Read.Shared',
      'Chat.Read',
      'Mail.Send',
      'https://graph.microsoft.com/User.Read',
    ],
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
