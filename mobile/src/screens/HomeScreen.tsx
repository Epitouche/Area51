import React, { useState, useContext, useEffect } from 'react';
import { Button, View, Text, TextInput, StyleSheet } from 'react-native';
import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParamList } from '../navigation/AppNavigator';
import LinearGradient from 'react-native-linear-gradient';
import { AppContext } from '../context/AppContext';
import { checkToken } from '../service/token';

type HomeScreenNavigationProp = StackNavigationProp<RootStackParamList, 'Home'>;

type Props = {
  navigation: HomeScreenNavigationProp;
};

export default function HomeScreen({ navigation }: Props) {
  const { serverIp, setServerIp } = useContext(AppContext);
  const [isConnected, setIsConnected] = useState(false);
  const [isValidIp, setIsValidIp] = useState(false);
  const [token, setToken] = useState('');

  const validateIp = (ip: string) => {
    const ipPattern =
      /^(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}$/;
    return ipPattern.test(ip);
  };

  const handleSave = () => {
    if (validateIp(serverIp)) {
      setIsValidIp(true);
      setIsConnected(true);
    } else {
      alert('Please enter a valid IP address');
      setIsValidIp(false);
      setIsConnected(false);
    }
  };

  useEffect(() => {
    checkToken('token', setToken);
    console.log('token:', token);
  },  []);

  return (
    <LinearGradient colors={['#7874FD', '#B225EE']} style={styles.wallpaper}>
      <View style={styles.container}>
        <Text>Home Screen</Text>
        <View style={styles.inputBox}>
          <TextInput
            style={styles.input}
            placeholder="Server IP"
            keyboardType="numeric"
            value={serverIp}
            onChangeText={setServerIp}
          />
          <Button title="Save" onPress={handleSave} />
        </View>
        {isValidIp && isConnected && (
          <View>
            <View style={styles.buttonBox}>
              <View style={styles.button}>
                <Button
                  title="Connexion"
                  onPress={() => navigation.navigate('Login')}
                />
              </View>
              <View style={styles.button}>
                <Button
                  title="Inscription"
                  onPress={() => navigation.navigate('Register')}
                />
              </View>
            </View>
            {token !== 'Error: token not found' && (
              <View style={styles.button}>
                <Button
                  title="Dashboard"
                  onPress={() => navigation.navigate('Dashboard')}
                />
              </View>
            )}
          </View>
        )}
      </View>
    </LinearGradient>
  );
}

const styles = StyleSheet.create({
  wallpaper: {
    flex: 1,
    alignItems: 'center',
    padding: '3%',
  },
  container: {
    width: '100%',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
  buttonBox: {
    width: '100%',
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
  button: {
    width: '45%',
  },
  inputBox: {
    width: '100%',
    flexDirection: 'row',
  },
  input: {
    backgroundColor: 'white',
    width: '80%',
    padding: 10,
    marginBottom: 10,
    borderRadius: 5,
  },
});
