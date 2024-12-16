import { getToken } from './getToken';

export const checkToken = async (
  tokenName: string,
  setToken: (token: string) => void,
) => {
  await getToken(tokenName, setToken);
};
