import React, { useState, useContext, useEffect } from 'react';
import { View, Text, TextInput, StyleSheet } from 'react-native';
import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParamList } from '../navigation/AppNavigator';
import { AppContext } from '../context/AppContext';
import { checkToken } from '../service/token';
import { getAboutJson } from '../service';
import { globalStyles } from '../styles/global_style';
import { Button } from 'react-native-paper';

type HomeScreenNavigationProp = StackNavigationProp<RootStackParamList, 'Home'>;

type Props = {
  navigation: HomeScreenNavigationProp;
};

export default function HomeScreen({ navigation }: Props) {
  const { serverIp, setServerIp, setAboutJson, isConnected } = useContext(AppContext);
  const [isValidIp, setIsValidIp] = useState(false);
  const [ipTmp, setIpTmp] = useState('');

  // const validateIp = (ip: string) => {
  //   const ipPattern =
  //     /^(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}$/;
  //   return ipPattern.test(ip);
  // };

  const handleSave = () => {
    // if (validateIp(serverIp)) {
      setIsValidIp(true);
      setServerIp(ipTmp);
    // } else {
    //   alert('Please enter a valid IP address');
    //   setIsValidIp(false);
    // }
  };

  useEffect(() => {
    if (isConnected) getAboutJson(serverIp, setAboutJson);
  }, [isConnected, serverIp]);

  return (
    <View style={globalStyles.wallpaper}>
      <View style={styles.container}>
        <Text style={globalStyles.titleWhite}>Area51</Text>
        <View style={styles.textAlign}>
          <View style={styles.textAlign}>
            <Text style={globalStyles.subtitleWhite}>Automate</Text>
          </View>
          <View>
            <Text style={globalStyles.subtitleWhite}>without limits</Text>
          </View>
        </View>
        <View style={styles.ipBox}>
          <TextInput
            style={[globalStyles.input, { width: '48%' }]}
            placeholder="Server IP"
            keyboardType="numeric"
            value={ipTmp}
            onChangeText={setIpTmp}
          />
          <Button
            onPress={handleSave}
            style={[globalStyles.buttonColor, styles.button]}>
            <Text style={globalStyles.textBlack}>Save</Text>
          </Button>
        </View>
        {isValidIp && (
          <View>
            <View style={styles.buttonBox}>
              <Button
                style={[globalStyles.buttonColor, styles.button]}
                onPress={() => navigation.navigate('Login')}>
                <Text style={globalStyles.textBlack}>Connexion</Text>
              </Button>
              <Button
                style={[globalStyles.buttonColor, styles.button]}
                onPress={() => navigation.navigate('Register')}>
                <Text style={globalStyles.textBlack}>Inscription</Text>
              </Button>
            </View>
            {isConnected && (
              <View style={styles.button}>
                <Button
                  style={[globalStyles.buttonColor, styles.button]}
                  onPress={() => navigation.navigate('Dashboard')}>
                  <Text style={globalStyles.textBlack}>Inscription</Text>
                </Button>
              </View>
            )}
          </View>
        )}
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    width: '100%',
    padding: '5%',
    paddingTop: '20%',
    alignItems: 'center',
    gap: 20,
  },
  buttonBox: {
    width: '100%',
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
  button: {
    width: '40%',
  },
  inputBox: {
    width: '100%',
    justifyContent: 'center',
    alignItems: 'center',
    flexDirection: 'row',
    gap: 10,
  },
  ipBox: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 20,
  },
  textAlign: {
    justifyContent: 'center',
    alignItems: 'center',
  },
});
