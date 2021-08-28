import Layout from '@/layout/index'
import componentsRouter from './components'
import chartsRouter from './charts'
import nestedRouter from './nested'
import tableRouter from './table'

const ExampleRouter = {
  path: '/my-game',
  component: Layout,
  redirect: '/my-game/index',
  name: 'example',
  meta: {
    title: '例子',
    icon: 'star'
  },
  children: [
    {
      path: '/permission',
      component: Layout,
      redirect: '/permission/page',
      alwaysShow: true, // will always show the root menu
      name: 'Permission',
      meta: {
        title: 'Permission',
        icon: 'lock',
        roles: ['admin', 'editor'] // you can set roles in root nav
      },
      children: [
        {
          path: '/permission',
          component: Layout,
          redirect: '/permission/page',
          alwaysShow: true, // will always show the root menu
          name: 'Permission',
          meta: {
            title: 'Permission',
            icon: 'lock',
            roles: ['admin', 'editor'] // you can set roles in root nav
          },
          children: [
            {
              path: 'page',
              component: () => import('@/views/ZExample/permission/page'),
              name: 'PagePermission',
              meta: {
                title: 'Page Permission',
                roles: ['admin'] // or you can only set roles in sub nav
              }
            },
            {
              path: 'directive',
              component: () => import('@/views/ZExample/permission/directive'),
              name: 'DirectivePermission',
              meta: {
                title: 'Directive Permission'
                // if do not set roles, means: this page does not require permission
              }
            },
            {
              path: 'role',
              component: () => import('@/views/ZExample/permission/role'),
              name: 'RolePermission',
              meta: {
                title: 'Role Permission',
                roles: ['admin']
              }
            }
          ]
        },

        {
          path: '/icon',
          component: Layout,
          children: [
            {
              path: 'index',
              component: () => import('@/views/ZExample/icons/index'),
              name: 'Icons',
              meta: { title: 'Icons', icon: 'icon', noCache: true }
            }
          ]
        },

        /** when your routing map is too long, you can split it into small modules **/
        componentsRouter,
        chartsRouter,
        nestedRouter,
        tableRouter,

        {
          path: '/example',
          component: Layout,
          redirect: '/example/list',
          name: 'Example',
          meta: {
            title: 'Example',
            icon: 'el-icon-s-help'
          },
          children: [
            {
              path: 'create',
              component: () => import('@/views/ZExample/example/create'),
              name: 'CreateArticle',
              meta: { title: 'Create Article', icon: 'edit' }
            },
            {
              path: 'edit/:id(\\d+)',
              component: () => import('@/views/ZExample/example/edit'),
              name: 'EditArticle',
              meta: { title: 'Edit Article', noCache: true, activeMenu: '/example/list' },
              hidden: true
            },
            {
              path: 'list',
              component: () => import('@/views/ZExample/example/list'),
              name: 'ArticleList',
              meta: { title: 'Article List', icon: 'list' }
            }
          ]
        },

        {
          path: '/tab',
          component: Layout,
          children: [
            {
              path: 'index',
              component: () => import('@/views/ZExample/tab/index'),
              name: 'Tab',
              meta: { title: 'Tab', icon: 'tab' }
            }
          ]
        },

        {
          path: '/error',
          component: Layout,
          redirect: 'noRedirect',
          name: 'ErrorPages',
          meta: {
            title: 'Error Pages',
            icon: '404'
          },
          children: [
            {
              path: '401',
              component: () => import('@/views/error-page/401'),
              name: 'Page401',
              meta: { title: '401', noCache: true }
            },
            {
              path: '404',
              component: () => import('@/views/error-page/404'),
              name: 'Page404',
              meta: { title: '404', noCache: true }
            }
          ]
        },

        {
          path: '/error-log',
          component: Layout,
          children: [
            {
              path: 'log',
              component: () => import('@/views/ZExample/error-log/index'),
              name: 'ErrorLog',
              meta: { title: 'Error Log', icon: 'bug' }
            }
          ]
        },

        {
          path: '/excel',
          component: Layout,
          redirect: '/excel/export-excel',
          name: 'Excel',
          meta: {
            title: 'Excel',
            icon: 'excel'
          },
          children: [
            {
              path: 'export-excel',
              component: () => import('@/views/ZExample/excel/export-excel'),
              name: 'ExportExcel',
              meta: { title: 'Export Excel' }
            },
            {
              path: 'export-selected-excel',
              component: () => import('@/views/ZExample/excel/select-excel'),
              name: 'SelectExcel',
              meta: { title: 'Export Selected' }
            },
            {
              path: 'export-merge-header',
              component: () => import('@/views/ZExample/excel/merge-header'),
              name: 'MergeHeader',
              meta: { title: 'Merge Header' }
            },
            {
              path: 'upload-excel',
              component: () => import('@/views/ZExample/excel/upload-excel'),
              name: 'UploadExcel',
              meta: { title: 'Upload Excel' }
            }
          ]
        },

        {
          path: '/zip',
          component: Layout,
          redirect: '/zip/download',
          alwaysShow: true,
          name: 'Zip',
          meta: { title: 'Zip', icon: 'zip' },
          children: [
            {
              path: 'download',
              component: () => import('@/views/ZExample/zip/index'),
              name: 'ExportZip',
              meta: { title: 'Export Zip' }
            }
          ]
        },

        {
          path: '/pdf',
          component: Layout,
          redirect: '/pdf/index',
          children: [
            {
              path: 'index',
              component: () => import('@/views/ZExample/pdf/index'),
              name: 'PDF',
              meta: { title: 'PDF', icon: 'pdf' }
            }
          ]
        },
        {
          path: '/pdf/download',
          component: () => import('@/views/ZExample/pdf/download'),
          hidden: true
        },

        {
          path: '/theme',
          component: Layout,
          children: [
            {
              path: 'index',
              component: () => import('@/views/ZExample/theme/index'),
              name: 'Theme',
              meta: { title: 'Theme', icon: 'theme' }
            }
          ]
        },

        {
          path: '/clipboard',
          component: Layout,
          children: [
            {
              path: 'index',
              component: () => import('@/views/ZExample/clipboard/index'),
              name: 'ClipboardDemo',
              meta: { title: 'Clipboard', icon: 'clipboard' }
            }
          ]
        },

        {
          path: 'external-link',
          component: Layout,
          children: [
            {
              path: 'https://github.com/PanJiaChen/vue-element-admin',
              meta: { title: 'External Link', icon: 'link' }
            }
          ]
        }]
    }
  ]
}

export default ExampleRouter
