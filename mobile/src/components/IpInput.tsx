import { StyleSheet, Text, TextInput, View } from 'react-native';
import { globalStyles } from '../styles/global_style';
import { Button } from 'react-native-paper';
import { useContext, useEffect, useState } from 'react';
import { AppContext } from '../context/AppContext';

export function IpInput() {
  const [ipTmp, setIpTmp] = useState('');
  const { setServerIp, serverIp, isBlackTheme } = useContext(AppContext);

  // const validateIp = (ip: string) => {
  //   const ipPattern =
  //     /^(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}$/;
  //   return ipPattern.test(ip);
  // };

  useEffect(() => {
    setIpTmp(serverIp);
  }, [serverIp]);

  const handleSave = () => {
    // if (validateIp(serverIp)) {
    setServerIp(ipTmp);
    // } else {
    //   alert('Please enter a valid IP address');
    //   setIsValidIp(false);
    // }
  };

  return (
    <View style={styles.ipBox}>
      <TextInput
        style={[
          isBlackTheme ? globalStyles.inputBlack : globalStyles.input,
          { width: '48%' },
        ]}
        placeholder="Server IP"
        keyboardType="numeric"
        value={ipTmp}
        onChangeText={setIpTmp}
      />
      <Button
        onPress={handleSave}
        style={[globalStyles.terciaryDark, styles.button]}>
        <Text style={isBlackTheme ? globalStyles.textBlack : globalStyles.text}>
          Save
        </Text>
      </Button>
    </View>
  );
}

const styles = StyleSheet.create({
  button: {
    width: '40%',
  },
  ipBox: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 20,
  },
});
