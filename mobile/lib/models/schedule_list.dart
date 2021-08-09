import 'package:holovn_mobile/models/schedule.dart';
import 'package:json_annotation/json_annotation.dart';
part 'schedule_list.g.dart';

@JsonSerializable(explicitToJson: true)
class ScheduleList{
  List<Schedule> values = [];
  ScheduleList(this.values);
  factory ScheduleList.fromJson(Map<String, dynamic> json) => _$ScheduleListFromJson(json);
}