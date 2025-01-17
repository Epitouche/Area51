import AsyncStorage from '@react-native-async-storage/async-storage';

export const saveToken = async (tokenName: string, token: string) => {
  try {
    await AsyncStorage.setItem(tokenName, token);
    console.log('Token stored');
  } catch (e) {
    console.error('Error storing the token', e);
  }
};
