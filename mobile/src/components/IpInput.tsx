import { StyleSheet, Text, TextInput, TouchableOpacity, View } from 'react-native';
import { globalStyles } from '../styles/global_style';
import { useContext, useEffect, useState } from 'react';
import { AppContext } from '../context/AppContext';
import { getToken, saveToken } from '../service';

export function IpInput() {
  const [ipTmp, setIpTmp] = useState('');
  const [checkIp, setcheckIp] = useState('');
  const { setServerIp, serverIp, isBlackTheme } = useContext(AppContext);

  // const validateIp = (ip: string) => {
  //   const ipPattern =
  //     /^(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}$/;
  //   return ipPattern.test(ip);
  // };

  useEffect(() => {
    const getIp = async () => {
      await getToken('serverIp', setcheckIp);
    };
    getIp();
  }, []);

  useEffect(() => {
    setIpTmp(serverIp);
  }, [serverIp]);

  const handleSave = () => {
    setServerIp(ipTmp);
    saveToken('serverIp', ipTmp);
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
      <TouchableOpacity
        onPress={handleSave}
        style={[globalStyles.buttonFormat, isBlackTheme ? globalStyles.primaryLight : globalStyles.primaryDark]}>
        <Text
          style={[
            isBlackTheme ? globalStyles.textColor : globalStyles.textColorBlack,
            globalStyles.textFormat,
            styles.button,
          ]}>
          Save
        </Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  button: {
    width: '100%',
  },
  ipBox: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 20,
  },
});
