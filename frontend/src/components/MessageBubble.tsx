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
        styles.wrapper,
        isMe ? styles.rightAlign : styles.leftAlign,
      ]}
    >
      <View style={[styles.bubble, isMe ? styles.rightBubble : styles.leftBubble]}>
        <Text style={[styles.text, isMe ? styles.whiteText : styles.blackText]}>{text}</Text>
      </View>
      <Text style={styles.time}>
        {timestamp.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
      </Text>
    </View>
  );
};

const styles = StyleSheet.create({
  wrapper: {
    margin: 5,
    maxWidth: '80%',
  },
  leftAlign: {
    alignSelf: 'flex-start',
    alignItems: 'flex-start',
  },
  rightAlign: {
    alignSelf: 'flex-end',
    alignItems: 'flex-end',
  },
  bubble: {
    padding: 10,
    borderRadius: 10,
  },
  leftBubble: {
    backgroundColor: '#D3D3D3',
    padding: 10,
    borderRadius: 10,
  },
  rightBubble: {
    backgroundColor: '#007AFF',
  },
  text: {
    fontSize: 16,
  },
  whiteText: {
    color: 'white',
  },
  blackText: {
    color: 'black',
  },
  time: {
    fontSize: 10,
    color: 'gray',
    marginTop: 2,
  },
});

export default MessageBubble;
