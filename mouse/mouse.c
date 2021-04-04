#include "windows.h"

void move(int x, int y)
{
    SetCursorPos(x, y);
}

void leftPress(int x, int y)
{
    SetCursorPos(x, y);
    mouse_event(MOUSEEVENTF_LEFTDOWN, x, y, 0, 0);
}

void leftRelease(int x, int y)
{
    SetCursorPos(x, y);
    mouse_event(MOUSEEVENTF_LEFTUP, x, y, 0, 0);
}

void rightPress(int x, int y)
{
    SetCursorPos(x, y);
    mouse_event(MOUSEEVENTF_RIGHTDOWN, x, y, 0, 0);
}

void rightRelease(int x, int y)
{
    SetCursorPos(x, y);
    mouse_event(MOUSEEVENTF_RIGHTUP, x, y, 0, 0);
}
