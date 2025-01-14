import { View, Text, StyleSheet, Image } from 'react-native';
import { globalStyles } from '../styles/global_style';
import { useContext, useState } from 'react';
import { AppContext } from '../context/AppContext';
import { Button } from 'react-native-paper';
import { getToken, parseServices, selectServicesParams } from '../service';
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
  useContext(AppContext);

  const handleOauthLogin = async () => {
    if (
      (await selectServicesParams({
        serverIp,
        serviceName: title,
        sessionToken: token,
      })) &&
      aboutJson
    )
      parseServices({
        aboutJson,
        serverIp,
        setServicesConnected,
      });
    else console.log('Failed');
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
      <Button
        onPress={handleOauthLogin}
        style={[
          styles.statusBar,
          status ? styles.connectedBar : styles.disconnectedBar,
        ]}>
        <Text style={styles.statusText}>
          {status ? 'Connected' : 'Disconnected'}
        </Text>
      </Button>
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
