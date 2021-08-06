// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'schedule.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Schedule _$ScheduleFromJson(Map<String, dynamic> json) {
  return Schedule(
    json['scheduleId'] as String,
    json['title'] as String,
    json['type'] as String,
    json['publishedAt'] as String,
    json['availableAt'] as String,
    json['duration'] as int,
    json['status'] as String,
    Channel.fromJson(json['channel'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$ScheduleToJson(Schedule instance) => <String, dynamic>{
      'scheduleId': instance.scheduleId,
      'title': instance.title,
      'type': instance.type,
      'publishedAt': instance.publishedAt,
      'availableAt': instance.availableAt,
      'duration': instance.duration,
      'status': instance.status,
      'channel': instance.channel.toJson(),
    };
