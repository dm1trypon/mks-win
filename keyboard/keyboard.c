#include "windows.h"

void keyDown(WORD code)
{
    const UINT mappedKey = MapVirtualKey(code, 0);

    INPUT input;
    input.type = INPUT_KEYBOARD;
    input.ki.dwFlags = KEYEVENTF_SCANCODE;
    input.ki.wScan = mappedKey;

    SendInput(1, &input, sizeof(input));
}

void keyUp(WORD code)
{
    const UINT mappedKey = MapVirtualKey(code, 0);

    INPUT input;
    input.type = INPUT_KEYBOARD;
    input.ki.dwFlags = KEYEVENTF_SCANCODE | KEYEVENTF_KEYUP;
    input.ki.wScan = mappedKey;
    
    SendInput(1, &input, sizeof(input));
}

short getKeyCode(char key)
{
    return VkKeyScan(key);   
}