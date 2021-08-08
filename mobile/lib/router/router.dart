import 'package:flutter/material.dart';
import 'package:holovn_mobile/models/schedule.dart';
import 'package:holovn_mobile/router/route_path.dart';
import 'package:holovn_mobile/screens/home.dart';
import 'package:holovn_mobile/screens/live.dart';

class AppRouter extends RouterDelegate<RoutePath>
    with ChangeNotifier, PopNavigatorRouterDelegateMixin<RoutePath> {
  final GlobalKey<NavigatorState> navigatorKey;
  final _pages = <Page>[MaterialPage(
    key: ValueKey("home"),
    child: HomePage(title: 'Holovn - A Vietnamese Hololive Fan App'),
  )];

  Schedule? _selectedSchedule;
  String? liveId;
  bool show404 = false;

  @override
  Widget build(BuildContext context) {
    return Navigator(
      key: navigatorKey,
      pages: List.of(_pages),
      onPopPage: (route, result) {
        if (!route.didPop(result)) {
          return false;
        }
        popRoute();
        return true;
      },
    );
  }

  @override
  Future<bool> popRoute() {
    if (_pages.length >= 1) {
      _pages.removeLast();
      notifyListeners();
      return Future.value(true);
    }
    return Future.value(false);
  }

  @override
  Future<void> setNewRoutePath(RoutePath path) async {
    if (path.isUnknown == null) {
      _selectedSchedule = null;
      show404 = true;
      return;
    }
    if (path.isLivePage) {
      if (path.schedule == null) {
        show404 = true;
        return;
      }
    }
    _selectedSchedule = null;
    show404 = false;
  }

  RoutePath get currentConfiguration {
    if (show404) {
      return RoutePath.unknown();
    }

    return _selectedSchedule == null && liveId == null
        ? RoutePath.home()
        : RoutePath.live(_selectedSchedule, liveId);
  }

  void navigate(Schedule? schedule, String? scheduleId) {
    _selectedSchedule = schedule;
    liveId = scheduleId;

    if (_selectedSchedule == null && liveId == null) {
      show404 = false;
      _pages.add(MaterialPage(
        key: ValueKey("home"),
        child: HomePage(title: 'Holovn - A Vietnamese Hololive Fan App'),
      ));
    } else if (_selectedSchedule != null || liveId != null) {
      _pages.add(MaterialPage(
        key: ValueKey("Live"),
        name: scheduleId,
        child: Live(schedule, liveId),
      ));
    }
    notifyListeners();
  }

  AppRouter() : navigatorKey = GlobalKey<NavigatorState>();
}
