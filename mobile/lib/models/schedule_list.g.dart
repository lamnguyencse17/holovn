// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'schedule_list.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ScheduleList _$ScheduleListFromJson(Map<String, dynamic> json) {
  return ScheduleList(
    (json['values'] as List<dynamic>)
        .map((e) => Schedule.fromJson(e as Map<String, dynamic>))
        .toList(),
  );
}

Map<String, dynamic> _$ScheduleListToJson(ScheduleList instance) =>
    <String, dynamic>{
      'values': instance.values.map((e) => e.toJson()).toList(),
    };
