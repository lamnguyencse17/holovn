// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'channel.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Channel _$ChannelFromJson(Map<String, dynamic> json) {
  return Channel(
    json['channelId'] as String,
    json['name'] as String,
    json['org'] as String,
    json['type'] as String,
    json['photo'] as String,
    json['englishName'] as String,
  );
}

Map<String, dynamic> _$ChannelToJson(Channel instance) => <String, dynamic>{
      'channelId': instance.channelId,
      'name': instance.name,
      'org': instance.org,
      'type': instance.type,
      'photo': instance.photo,
      'englishName': instance.englishName,
    };
