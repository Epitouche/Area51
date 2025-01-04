import { getToken } from './getToken';

export async function checkToken(tokenName: string): Promise<boolean> {
  let token = '';
  const setToken = (newToken: string) => {
    token = newToken;
  };

  if ((await getToken(tokenName, setToken)) === false) return false;

  if (token === 'Error: getting the token') return false;
  return true;
}
