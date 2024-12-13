import { useState } from 'react';
import { getToken } from './getToken';

export const checkToken = async (
  tokenName: string,
) => {
  const [token, setToken] = useState('');
  await getToken(tokenName, setToken);

  if (token === 'Error: getting the token')
    return false;
  return true
};
