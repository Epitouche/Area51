import { View, Text, StyleSheet, Image, TouchableOpacity } from 'react-native';
import { globalStyles } from '../styles/global_style';
import { selectServicesParams } from '../service';
import { AboutJson } from '../types';
import { useEffect, useState } from 'react';

interface ServiceCardProps {
  title: string;
  image: string;
  status: boolean;
  isBlackTheme?: boolean;
  aboutJson: AboutJson;
  serverIp: string;
  token: string;
  oauth: boolean;
  setNeedRefresh: (needRefresh: boolean) => void;
  setModalVisible: (modalvisible: boolean) => void;
  setSelectedService: (selectedServices: string) => void;
}

export function ServiceCard({
  image,
  status,
  title,
  isBlackTheme,
  serverIp,
  token,
  oauth,
  setNeedRefresh,
  setModalVisible,
  setSelectedService,
}: ServiceCardProps) {
  const [isConnected, setIsConnected] = useState(status);

  useEffect(() => {
    if (!oauth) setIsConnected(true);
    else setIsConnected(status);
  }, [status]);

  const handleOauthLogin = async (
    isConnected: boolean,
    serviceName: string,
  ) => {
    if (!oauth) return;
    if (isConnected) {
      setModalVisible(true);
      setSelectedService(serviceName);
    } else {
      await selectServicesParams({
        serverIp,
        serviceName: serviceName,
        sessionToken: token,
      });
      setNeedRefresh(true);
    }
  };

  return (
    <View
      style={[
        isBlackTheme ? globalStyles.primaryLight : globalStyles.terciaryDark,
        styles.card,
        status ? styles.connected : styles.disconnected,
      ]}>
      <Image
        source={{
          uri: image,
        }}
        style={styles.logo}
      />
      <Text
        style={[
          isBlackTheme ? globalStyles.textColor : globalStyles.textColorBlack,
          styles.title,
        ]}>
        {title[0].toLocaleUpperCase() + title.slice(1)}
      </Text>
      <TouchableOpacity
        onPress={() => handleOauthLogin(status, title)}
        disabled={!oauth}
        style={[
          styles.statusBar,
          isConnected ? styles.connectedBar : styles.disconnectedBar,
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
