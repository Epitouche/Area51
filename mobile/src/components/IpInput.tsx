import {
  StyleSheet,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from 'react-native';
import { globalStyles } from '../styles/global_style';
import { useEffect, useState } from 'react';
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
            isBlackTheme ? globalStyles.textColor : globalStyles.textColorBlack,
          ]}
          accessibilityLabel="Bullet">
          •
        </Text>
        <Text
          style={[
            isBlackTheme ? globalStyles.textColor : globalStyles.textColorBlack,
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
            { width: '50%' },
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
            isBlackTheme
              ? globalStyles.secondaryDark
              : globalStyles.primaryLight,
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
          isBlackTheme ? globalStyles.secondaryDark : globalStyles.primaryLight,
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
    padding: 20,
    borderRadius: 10,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.8,
    shadowRadius: 2,
    elevation: 5,
    margin: 20,
  },
  button: {
    width: '100%',
  },
  ipBox: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 20,
    width: '100%',
    marginBottom: 20,
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
