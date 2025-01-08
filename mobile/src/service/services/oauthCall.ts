import { AuthConfiguration, authorize } from 'react-native-app-auth';

type AuthApiCall = {
  config: AuthConfiguration;
  setToken: (accessToken: string) => void;
};

export async function OauthLogin({ config, setToken }: AuthApiCall) {
  const authState = await authorize(config);
  if (!authState.accessToken) return false;
  setToken(authState.accessToken);
  return true;
}
