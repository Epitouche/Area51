import {
  StyleSheet,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from 'react-native';
import { globalStyles } from '../styles/global_style';
import { useContext, useEffect, useState } from 'react';
import { AppContext } from '../context/AppContext';
import { getToken, refreshServices, saveToken } from '../service';
import { AboutJson, AboutJsonParse } from '../types';
interface IpInputProps {
  isBlackTheme?: boolean;
  setAboutJson: (aboutJson: AboutJson) => void;
  setServicesConnected: (servicesConnected: AboutJsonParse) => void;
  aboutJson: AboutJson | undefined;
  setServerIp: (serverIp: string) => void;
  serverIp: string;
}

export function IpInput({
  aboutJson,
  setAboutJson,
  setServicesConnected,
  serverIp,
  setServerIp,
  isBlackTheme,
}: IpInputProps) {
  const [ipTmp, setIpTmp] = useState('');
  const [checkIp, setcheckIp] = useState('');

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
    <View
      style={[
        isBlackTheme ? globalStyles.primaryLight : globalStyles.terciaryLight,
        styles.card,
      ]}>
        <View style={{ flexDirection: 'row', alignItems: 'center' }}>
          <Text
            style={[
              styles.bullet,
              isBlackTheme
                ? globalStyles.textColor
                : globalStyles.textColorBlack,
            ]}
            accessibilityLabel="Bullet">
            •
          </Text>
          <Text
            style={[
              isBlackTheme
                ? globalStyles.textColor
                : globalStyles.textColorBlack,
              styles.subtitle,
            ]}
            accessibilityLabel="Set a server Ip">
            Set a server Ip
          </Text>
        </View>
      <View style={styles.ipBox}>
        <TextInput
          style={[
            isBlackTheme ? globalStyles.input : globalStyles.inputBlack,
            { width: '48%' },
          ]}
          placeholder="Server IP"
          keyboardType="numeric"
          value={ipTmp}
          onChangeText={setIpTmp}
        />
        <TouchableOpacity
          onPress={handleSave}
          style={[
            globalStyles.buttonFormat,
            isBlackTheme ? globalStyles.primaryDark : globalStyles.primaryLight,
          ]}>
          <Text
            style={[
              isBlackTheme
                ? globalStyles.textColorBlack
                : globalStyles.textColor,
              globalStyles.textFormat,
              styles.button,
            ]}>
            Save
          </Text>
        </TouchableOpacity>
      </View>
      <TouchableOpacity
        style={[
          globalStyles.buttonFormat,
          isBlackTheme ? globalStyles.primaryDark : globalStyles.primaryLight,
        ]}
        onPress={() =>
          refreshServices({
            serverIp,
            setAboutJson,
            setServicesConnected,
            aboutJson,
          })
        }>
        <Text
          style={[
            isBlackTheme ? globalStyles.textColorBlack : globalStyles.textColor,
            globalStyles.textFormat,
          ]}>
          Refresh
        </Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  card: {
    width: '100%',
    borderRadius: 10,
    alignItems: 'center',
    gap: 10,
    paddingBottom: 20,
    paddingTop: 20,
  },
  button: {
    width: '100%',
  },
  ipBox: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 20,
  },
  bullet: {
    fontSize: 20,
    marginRight: 10,
  },
  subtitle: {
    fontSize: 16,
    fontWeight: '600',
  },
});
