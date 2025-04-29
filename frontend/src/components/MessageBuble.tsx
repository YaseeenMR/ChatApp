import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

interface MessageBubbleProps {
  text: string;
  isMe: boolean;
  timestamp: Date;
}

const MessageBubble = ({ text, isMe, timestamp }: MessageBubbleProps) => {
  return (
    <View
      style={[
        styles.container,
        isMe ? styles.rightContainer : styles.leftContainer,
      ]}
    >
      <Text style={styles.text}>{text}</Text>
      <Text style={styles.time}>
        {timestamp.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
      </Text>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    padding: 10,
    margin: 5,
    borderRadius: 10,
    maxWidth: '80%',
  },
  leftContainer: {
    backgroundColor: '#e5e5ea',
    alignSelf: 'flex-start',
  },
  rightContainer: {
    backgroundColor: '#007AFF',
    alignSelf: 'flex-end',
  },
  text: {
    color: 'black',
  },
  time: {
    fontSize: 10,
    color: 'gray',
    textAlign: 'right',
    marginTop: 5,
  },
});

export default MessageBubble;