import 'package:holovn_mobile/models/schedule.dart';

class RoutePath {
  final Schedule? schedule;
  final bool? isUnknown;
  final String? liveId;

  RoutePath.home()
      : schedule = null,
        liveId = null,
        isUnknown = false;

  RoutePath.live(this.schedule, this.liveId) : isUnknown = false;

  RoutePath.unknown()
      : schedule = null,
        liveId = null,
        isUnknown = true;

  bool get isHomePage => schedule == null && liveId == null;

  bool get isLivePage => schedule != null || liveId != null;
}
