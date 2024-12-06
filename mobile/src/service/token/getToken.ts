import AsyncStorage from '@react-native-async-storage/async-storage';

export const useGetToken = async (setToken: (Token: string) => void) => {
  try {
    const token = await AsyncStorage.getItem('@user_token');
    token ? setToken(token) : setToken('Error: token not found');
    return true;
  } catch (e) {
    console.error('Error getting the token', e);
    setToken('Error: getting the token');
    return false;
  }
};
