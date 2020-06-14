#ifndef _OBJECT_MANAGER_H_
#define _OBJECT_MANAGER_H_


#ifndef HU_RET_OK
#define HU_RET_OK 0
#define HU_RET_PTR_ERROR 1
#define HU_RET_FILE_ERROR 2
#define HU_RET_CONTENT_ERROR 3
#define HU_RET_PARAMS_ERROR 4
#define HU_RET_TYPE_ERROR 5
#define HU_RET_ERROR 6
#define HU_RET_STATE_ERROR 7
#endif


#if defined(ANDROID)
#include <jni.h>
#include <android/log.h>
#define LIB_EXPORT JNIEXPORT
#elif defined(WIN32)
#define LIB_EXPORT __declspec(dllexport)
#else
#define LIB_EXPORT
#endif

typedef void HuObjManager;
typedef struct {
    int x;
    int y;
    int width;
    int height;
} HRect;

/***********************************************************
 * Load configuration files and create object manager
 * params:
 *      modelDir   configuration files directory
 *      tag        business tag
 *      pmanager   the pointer of object manager
 *                 which used to recognise objects
 *
 * return:
 *      0        success
 *      other    failed
 ***********************************************************/
LIB_EXPORT int load_object_manager(const char *modelDir, const char *tag, HuObjManager **pmanager);


/***********************************************************
 * Release object manager
 * params:
 *      pmanager  the pointer of object manager
 ***********************************************************/
LIB_EXPORT void release_object_manager(HuObjManager **pmanager);


/************************************************************
 * Detect objects in image
 * params:
 *      manager     object manager
 *      imgPath     path of image
 *      srect       candidate rect
 *      rrects      result rectangles, malloc inner
 *      rsize       result size
 *
 * return:
 *      0 success, other failed
 *
 ************************************************************/
LIB_EXPORT int detect_objects(void *ptr, const char *imgPath, HRect srect, HRect **rrects, int *rsize);


#endif

