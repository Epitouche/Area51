import AsyncStorage from '@react-native-async-storage/async-storage';

export const deleteToken = async (tokenName: string) => {
  try {
    await AsyncStorage.removeItem(tokenName);
    return true;
  } catch (e) {
    console.error('Error removing the token', e);
    return false;
  }
};
