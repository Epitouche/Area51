import { Image, StyleSheet, Text, View } from 'react-native';
import { Button } from 'react-native-paper';

export interface OauthLoginButtonProps {
  handleOauthLogin: () => void;
  color?: string;
  img: string;
  name: string;
}

// https://img.icons8.com/?size=100&id=12599&format=png

 export function OauthLoginButton({
   handleOauthLogin,
   color,
   name,
   img,
 }: OauthLoginButtonProps) {
   if (!color) {
     color = '#FFFFFF';
   }
   return (
     <Button
       onPress={handleOauthLogin}
       style={[styles.button, { backgroundColor: color }]}>
       <View style={styles.buttonContent}>
         <Image source={{ uri: img }} style={styles.icon} />
         <Text style={styles.text}>{name}</Text>
       </View>
     </Button>
   );
 }

const styles = StyleSheet.create({
  button: {
    width: 'auto',
    marginTop: 10,
    marginBottom: 10,
    alignItems: 'center',
    flexDirection: 'row',
    justifyContent: 'center',
  },
  buttonContent: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  icon: {
    width: 20,
    height: 20,
    marginRight: 10,
  },
  text: {
    color: '#222831',
    fontSize: 16,
    fontWeight: 'bold',
  },
});
