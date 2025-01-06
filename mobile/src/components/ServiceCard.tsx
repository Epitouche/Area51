import { View, Text, StyleSheet, Image } from 'react-native';
import { globalStyles } from '../styles/global_style';
import { useContext } from 'react';
import { AppContext } from '../context/AppContext';
import { Button } from 'react-native-paper';

interface ServiceCardProps {
  title: string;
  image: string;
  status: string;
  handleOauthLogin: () => void;
}

export function ServiceCard({ image, status, title, handleOauthLogin }: ServiceCardProps) {
  const { isBlackTheme } = useContext(AppContext);
  return (
    <View
      style={[
        styles.card,
        status === 'Connected' ? styles.connected : styles.disconnected,
      ]}>
      <Image
        source={{
          uri: image,
        }}
        style={styles.logo}
      />
      <Text
        style={[
          isBlackTheme ? globalStyles.text : globalStyles.textBlack,
          styles.title,
        ]}>
        {title[0].toLocaleUpperCase() + title.slice(1)}
      </Text>
      <Button
        onPress={handleOauthLogin}
        style={[
          styles.statusBar,
          status === 'Connected' ? styles.connectedBar : styles.disconnectedBar,
        ]}>
        <Text style={styles.statusText}>{status}</Text>
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
