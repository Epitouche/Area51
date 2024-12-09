import AsyncStorage from '@react-native-async-storage/async-storage';

export const saveToken = async (tokenName: string, token: string) => {
  try {
    await AsyncStorage.setItem(tokenName, token);
  } catch (e) {
    console.error('Error storing the token', e);
  }
};
