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
    DateTime.parse(json['publishedAt'] as String),
    DateTime.parse(json['availableAt'] as String),
    DateTime.parse(json['startScheduled'] as String),
    json['duration'] as int,
    json['status'] as String,
    Channel.fromJson(json['channel'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$ScheduleToJson(Schedule instance) => <String, dynamic>{
      'scheduleId': instance.scheduleId,
      'title': instance.title,
      'type': instance.type,
      'publishedAt': instance.publishedAt.toIso8601String(),
      'availableAt': instance.availableAt.toIso8601String(),
      'startScheduled': instance.startScheduled.toIso8601String(),
      'duration': instance.duration,
      'status': instance.status,
      'channel': instance.channel.toJson(),
    };
