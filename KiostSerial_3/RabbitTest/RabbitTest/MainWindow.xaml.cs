using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

using System.Windows.Threading;

using RabbitMQ.Client;

//Using namespaces 
using System.Data;
using MySql.Data.MySqlClient;
using System.Configuration;
namespace RabbitTest
{
    /// <summary>
    /// MainWindow.xaml에 대한 상호 작용 논리
    /// </summary>
    public partial class MainWindow : Window
    {
        Receive receive;

        int im_cnt = 0;

        public MainWindow()
        {
            InitializeComponent();
            receive = new Receive(this);
            System.Threading.Timer timer = new System.Threading.Timer(CallBack);
            Send send = new Send("localhost", "hello",  "test");
            timer.Change(0, 1000);
        }

        delegate void TimerEventFiredDelegate();
        void CallBack(object state)
        {
            Dispatcher.BeginInvoke(new TimerEventFiredDelegate(Work));
        }

        private void Work()
        {
        //수행해야할 작업(UI Thread 핸들링 가능)
            im_cnt++;
            sendTxt.Text = im_cnt.ToString();
            receiveTxt.Text = receive.msg1;
            Send send = new Send("localhost", "hello", sendTxt.Text+"ddd");
        }


        private void sendBtn_Click(object sender, RoutedEventArgs e)
        {
            //           Send send = new Send("localhost", "hello", "cmd,"+ sendTxt.Text);
            {

                //port.Write(Command.InitializePrinter);

                string strmsg = "cmd,";
                strmsg += ConvertFontSize(6, 6);
                strmsg += "123호";
                strmsg += Command.NewLine;
                strmsg += Command.LineFeed(10);
                strmsg += ConvertFontSize(1, 1);
                strmsg += "123호입니다 .";
                strmsg += Command.NewLine;
                strmsg += "----------------";
                strmsg += Command.NewLine;
                strmsg += Command.LineFeed(10);
                strmsg += Command.Cut;

                //System.Text.Encoding.GetEncoding(51949).GetBytes(strmsg);
                //Send send = new Send("localhost", "hello", Encoding.UTF8.GetString(System.Text.Encoding.GetEncoding("ks_c_5601-1987").GetBytes(strmsg)));
                Send send = new Send("localhost", "hello", Encoding.UTF8.GetString(System.Text.Encoding.GetEncoding(51949).GetBytes(strmsg)));

            }
        }


        private void Window_Closing(object sender, System.ComponentModel.CancelEventArgs e)
        {
            receive.close();
        }
        /// <summary>
        /// FONT 명령어의 글자사이즈 속성을 변환 합니다.
        /// </summary>
        /// <param name="width">가로</param>
        /// <param name="height">세로</param>
        /// <returns>가로 x 세로</returns>
        public string ConvertFontSize(int width, int height)
        {
            string result = "0";
            int _w, _h;

            //가로변환
            if (width == 1)
                _w = 0;
            else if (width == 2)
                _w = 16;
            else if (width == 3)
                _w = 32;
            else if (width == 4)
                _w = 48;
            else if (width == 5)
                _w = 64;
            else if (width == 6)
                _w = 80;
            else if (width == 7)
                _w = 96;
            else if (width == 8)
                _w = 112;
            else _w = 0;

            //세로변환
            if (height == 1)
                _h = 0;
            else if (height == 2)
                _h = 1;
            else if (height == 3)
                _h = 2;
            else if (height == 4)
                _h = 3;
            else if (height == 5)
                _h = 4;
            else if (height == 6)
                _h = 5;
            else if (height == 7)
                _h = 6;
            else if (height == 8)
                _h = 7;
            else _h = 0;

            //가로x세로
            int sum = _w + _h;
            result = Command.GS + "!" + Command.DecimalToCharString(sum);

            return result;
        }
    }
}
