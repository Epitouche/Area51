import { View, Text, StyleSheet, Image, TouchableOpacity } from 'react-native';
import { globalStyles } from '../styles/global_style';
import {
  parseServices,
  selectServicesParams,
  logoutServices,
} from '../service';
import { AboutJson, AboutJsonParse } from '../types';

interface ServiceCardProps {
  title: string;
  image: string;
  status: boolean;
  isMobile?: boolean;
  setServicesConnected: (services: AboutJsonParse) => void;
  aboutJson: AboutJson;
  serverIp: string;
  token: string;
}

export function ServiceCard({
  image,
  status,
  title,
  isMobile,
  aboutJson,
  serverIp,
  token,
  setServicesConnected,
}: ServiceCardProps) {

  const handleOauthLogin = async (
    isConnected: boolean,
    serviceName: string,
  ) => {
    if (isConnected) {
      await logoutServices(serverIp, serviceName, token);
    } else {
      await selectServicesParams({
        serverIp,
        serviceName: serviceName,
        sessionToken: token,
      });
    }
    if (aboutJson)
      await parseServices({
        aboutJson,
        serverIp,
        setServicesConnected,
      });
  };

  return (
    <View
      style={[styles.card, status ? styles.connected : styles.disconnected]}>
      <Image
        source={{
          uri: image,
        }}
        style={styles.logo}
      />
      <Text
        style={[
          isMobile ? globalStyles.textColor : globalStyles.textColorBlack,
          styles.title,
        ]}>
        {title[0].toLocaleUpperCase() + title.slice(1)}
      </Text>
      <TouchableOpacity
        onPress={() => handleOauthLogin(status, title)}
        style={[
          styles.statusBar,
          status ? styles.connectedBar : styles.disconnectedBar,
        ]}>
        <Text style={styles.statusText}>
          {status ? 'Connected' : 'Disconnected'}
        </Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  card: {
    width: 110,
    borderRadius: 10,
    backgroundColor: '#f0f0f0',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  logo: {
    marginTop: 10,
    width: 40,
    height: 40,
  },
  title: {
    fontSize: 16,
    fontWeight: 'bold',
    marginBottom: 10,
  },
  statusBar: {
    width: '100%',
    alignItems: 'center',
    justifyContent: 'center',
    borderTopLeftRadius: 0,
    borderTopRightRadius: 0,
    borderBottomLeftRadius: 10,
    borderBottomRightRadius: 10,
    padding: 8,
  },
  statusText: {
    fontSize: 12,
    color: '#fff',
    fontWeight: 'bold',
  },
  connected: {
    borderColor: '#28a745',
  },
  disconnected: {
    borderColor: '#6c757d',
  },
  connectedBar: {
    backgroundColor: '#28a745',
  },
  disconnectedBar: {
    backgroundColor: '#6c757d',
  },
});
