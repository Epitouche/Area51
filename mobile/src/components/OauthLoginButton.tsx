import { Image, StyleSheet, Text, View, TouchableOpacity } from 'react-native';
import { selectServicesParams } from '../service';
import { globalStyles } from '../styles/global_style';

export interface OauthLoginButtonProps {
  img?: string;
  name: string;
  serverIp: string;
  setIsConnected: (isConnected: boolean) => void;
  isBlackTheme?: boolean;
}

export function OauthLoginButton({
  name,
  img,
  serverIp,
  setIsConnected,
  isBlackTheme,
}: OauthLoginButtonProps) {

  const handleOauthLogin = async () => {
    if (await selectServicesParams({ serverIp, serviceName: name }) === true) {
      setIsConnected(true);
    }
  };

  return (
    <TouchableOpacity
      onPress={handleOauthLogin}
      style={[
        globalStyles.buttonFormat,
        isBlackTheme ? globalStyles.secondaryLight : globalStyles.terciaryLight,
      ]}>
      <View style={styles.buttonContent}>
        <Image source={{ uri: img }} style={styles.icon} />
        <Text
          style={[
            globalStyles.textFormat,
            isBlackTheme ? globalStyles.textColor : globalStyles.textColorBlack,
          ]}>
          {name.charAt(0).toUpperCase() + name.slice(1)}
        </Text>
      </View>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  buttonContent: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  icon: {
    width: 20,
    height: 20,
    marginRight: 10,
  },
});
