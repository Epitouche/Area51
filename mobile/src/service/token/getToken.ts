import AsyncStorage from '@react-native-async-storage/async-storage';

export async function getToken(
  tokenName: string,
  setToken: (Token: string) => void,
) {
  try {
    const token = await AsyncStorage.getItem(tokenName);
    token ? setToken(token) : setToken('Error: token not found');
    return true;
  } catch (e) {
    console.error('Error getting the token', e);
    setToken('Error: token not found');
    return false;
  }
}
